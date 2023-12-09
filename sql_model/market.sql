

CREATE TABLE "category" (
    "id" UUID NOT NULL PRIMARY KEY,
    "title" VARCHAR(46) NOT NULL,
    "parent_id" UUID REFERENCES "category" ("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE "product" (
    "id" UUID NOT NULL PRIMARY KEY,
    "name" VARCHAR(46) NOT NULL,
    "barcode" VARCHAR NOT NULL,
    "price" NUMERIC NOT NULL,
    "image_url" VARCHAR(255) NOT NULL,
    "category_id" UUID NOT NULL REFERENCES "category"("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE "user" (
"id" UUID PRIMARY KEY,
"name" VARCHAR(255) NOT NULL,
"logi" VARCHAR(255) UNIQUE NOT NULL,
"passwor" VARCHAR(255) NOT NULL,
"expired_at" TIMESTAMP,
"created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
"updated_at" TIMESTAMP
);
