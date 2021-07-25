
CREATE TABLE "users" (
  "id" serial PRIMARY KEY,
  "email" text NOT NULL,
  "password" text NOT NULL,
  "created_at" timestamptz
);

CREATE TABLE "groups" (
  "id" serial PRIMARY KEY,
  "organizer" integer NOT NULL,
  "name" text NOT NULL,
  "description" text,
  "created_at" timestamptz
);

CREATE TABLE "group_books" (
  "id" serial PRIMARY KEY,
  "group_id" serial,
  "book_id" serial,
  "completion" float NOT NULL,
  "last_modified" timestamptz
);

CREATE TABLE "events" (
  "id" serial PRIMARY KEY,
  "group_id" integer,
  "book_id" integer,
  "chapter_id" integer,
  "video_link" text NOT NULL,
  "start_time" timestamptz NOT NULL,
  "duration" float,
  "description" text,
  "canceled" boolean,
  "created_at" timestamptz
);

CREATE TABLE "event_attendees" (
  "id" serial PRIMARY KEY,
  "event_id" integer,
  "user_id" integer
);

CREATE TABLE "event_sections" (
  "id" serial PRIMARY KEY,
  "event_id" integer,
  "section_id" integer,
  "presenter" integer,
  "complete" boolean
);

CREATE TABLE "books" (
  "id" serial PRIMARY KEY,
  "title" text NOT NULL,
  "author" text NOT NULL,
  "category" text NOT NULL
);

CREATE TABLE "chapters" (
  "id" serial PRIMARY KEY,
  "book_id" integer,
  "title" text NOT NULL,
  "number" integer
);

CREATE TABLE "sections" (
  "id" serial PRIMARY KEY,
  "chapter_id" integer,
  "title" text,
  "number" integer
);

ALTER TABLE "groups" ADD FOREIGN KEY ("organizer") REFERENCES "users" ("id");

ALTER TABLE "group_books" ADD FOREIGN KEY ("group_id") REFERENCES "groups" ("id");

ALTER TABLE "group_books" ADD FOREIGN KEY ("book_id") REFERENCES "books" ("id");

ALTER TABLE "events" ADD FOREIGN KEY ("group_id") REFERENCES "groups" ("id");

ALTER TABLE "events" ADD FOREIGN KEY ("book_id") REFERENCES "books" ("id");

ALTER TABLE "events" ADD FOREIGN KEY ("chapter_id") REFERENCES "chapters" ("id");

ALTER TABLE "event_attendees" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id");

ALTER TABLE "event_attendees" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "event_sections" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id");

ALTER TABLE "event_sections" ADD FOREIGN KEY ("section_id") REFERENCES "sections" ("id");

ALTER TABLE "event_sections" ADD FOREIGN KEY ("presenter") REFERENCES "users" ("id");

ALTER TABLE "chapters" ADD FOREIGN KEY ("book_id") REFERENCES "books" ("id");

ALTER TABLE "sections" ADD FOREIGN KEY ("chapter_id") REFERENCES "chapters" ("id");

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "events" ("book_id");

CREATE INDEX ON "events" ("start_time");

CREATE INDEX ON "books" ("category");

CREATE INDEX ON "chapters" ("book_id");

CREATE INDEX ON "chapters" ("number");

CREATE INDEX ON "sections" ("chapter_id");

CREATE INDEX ON "sections" ("number");