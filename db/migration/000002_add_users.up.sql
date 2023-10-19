CREATE TABLE "users" (
                         "username" varchar PRIMARY KEY ,
                         "hashed_password" varchar NOT NULL ,
                         "full_name" varchar NOT NULL ,
                         "email" varchar NOT NULL ,
                         "password_changed_at" timestamptz NOT NULL DEFAULT('01-01-0001 00:00:00+00') ,
                         "created_at" timestamptz NOT NULL DEFAULT(now())
);

ALTER TABLE "accounts" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

--CREATE UNIQUE INDEX ON "accounts" ("owner","currency");
ALTER TABLE "accounts" ADD CONSTRAINT "owner_currency_key" UNIQUE ("owner", "currency");
