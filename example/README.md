The code in this directory was generated with the solution in this repo, against a db defined with the following sql:

```
                List of relations
 Schema |      Name      |   Type   |   Owner
--------+----------------+----------+------------
 public | authors        | table    | gnorm-user
 public | books          | table    | gnorm-user
(4 rows)


               Table "public.authors"
 Column | Type |              Modifiers
--------+------+-------------------------------------
 id     | uuid | not null default uuid_generate_v4()
 name   | text | not null
Indexes:
    "authors_pkey" PRIMARY KEY, btree (id)
    "authors_name_idx" btree (name)
Referenced by:
    TABLE "books" CONSTRAINT "books_author_id_fkey" FOREIGN KEY (author_id) REFERENCES authors(id)


                                               Table "public.books"
  Column   |           Type           |                                 Modifiers
-----------+--------------------------+----------------------------------------------------------------------------
 id        | integer                  | not null default nextval('books_id_seq'::regclass)
 author_id | uuid                     | not null
 isbn      | character(32)            | not null
 booktype  | book_type                | not null
 title     | text                     | not null
 pages     | integer                  | not null
 summary   | text                     |
 available | timestamp with time zone | not null default '2017-09-04 21:43:39.197538-04'::timestamp with time zone
Indexes:
    "books_pkey" PRIMARY KEY, btree (id)
    "books_isbn_key" UNIQUE CONSTRAINT, btree (isbn)
    "books_title_idx" btree (author_id, title)
Foreign-key constraints:
    "books_author_id_fkey" FOREIGN KEY (author_id) REFERENCES authors(id)

    
                                          List of data types
 Schema |   Name    | Internal name | Size |  Elements  |   Owner    | Access privileges | Description
--------+-----------+---------------+------+------------+------------+-------------------+-------------
 public | book_type | book_type     | 4    | FICTION   +| gnorm-user |                   |
        |           |               |      | NONFICTION |            |                   |
    
```