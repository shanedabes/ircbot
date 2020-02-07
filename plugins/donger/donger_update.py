#!/usr/bin/env python

import urllib.request
import json

DONGER_URL_TMPL = (
    "http://dongerlist.com/wp-json/wp/v2/posts?per_page=100&page={page_no}"
)
PAGE_COUNT = 12
DONGER_CAT_URL_TMPL = (
    "http://dongerlist.com/wp-json/wp/v2/categories?per_page=100&page={page_no}"
)
CAT_PAGE_COUNT = 2


def get_page(tmpl, page_no):
    url = tmpl.format(page_no=page_no)
    r = urllib.request.urlopen(url).read()

    return r


def page_to_json(page):
    j = json.loads(page)

    def e(donger, categories):
        return {
            "donger": donger,
            "categories": categories
        }

    return [e(i["title"]["rendered"], i["categories"]) for i in j]


def cat_page_to_json(page):
    j = json.loads(page)

    return {i["id"]: i["name"].lower() for i in j}


def donger_name_categories(donger, categories):
    categories_named = [categories[i] for i in donger["categories"]]

    donger["categories"] = categories_named
    return donger


def main():
    print("package donger")
    print()
    print("var dongerCollection = dongers{")

    pages = [get_page(DONGER_URL_TMPL, i) for i in range(1, PAGE_COUNT+1)]
    dongers = [i for s in pages for i in page_to_json(s)]

    cat_pages = [get_page(DONGER_CAT_URL_TMPL, i) for i in range(1, CAT_PAGE_COUNT+1)]
    cats = {k: v for p in cat_pages for k, v in cat_page_to_json(p).items()}

    dongers2 = [donger_name_categories(d, cats) for d in dongers]

    donger_cats = {c: [d["donger"] for d in dongers2 if c in d["categories"]] for c in cats.values()}

    for k in donger_cats:
        print('	"{}": []string{{'.format(k))
        for d in donger_cats[k]:
            print(u'		"{}",'.format(d))
        print('	},')

    print("}")


if __name__ == "__main__":
    main()
