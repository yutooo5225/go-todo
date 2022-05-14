CREATE TABLE "task" (
  "task_id" int PRIMARY KEY,
  "list_id" int,
  "discription" varchar,
  "done" boolean,
  "create_at" timestamp,
  "update_at" timestamp
);

CREATE TABLE "list" (
  "list_id" int PRIMARY KEY,
  "list_name" varchar
);

ALTER TABLE "task" ADD FOREIGN KEY ("list_id") REFERENCES "list" ("list_id");
