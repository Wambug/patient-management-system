CREATE TABLE "physician" (
  "doctorid" SERIAL PRIMARY KEY,
  "username" varchar UNIQUE,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01--01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "patient" (
  "patientid" SERIAL PRIMARY KEY,
  "username" varchar UNIQUE,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "dob" timestamp NOT NULL,
  "contact" varchar UNIQUE NOT NULL,
  "bloodgroup" varchar NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01--01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "patientrecords" (
  "recordid" SERIAL PRIMARY KEY,
  "patientid" integer,
  "date" timestamp NOT NULL,
  "disease" varchar NOT NULL,
  "prescription" varchar NOT NULL,
  "diagnosis" varchar NOT NULL,
  "weight" varchar NOT NULL,
  "doctorid" integer NOT NULL
);

CREATE TABLE "appointment" (
  "appointmentid" SERIAL PRIMARY KEY,
  "doctorid" integer NOT NULL,
  "patientid" integer NOT NULL,
  "appointmentdate" timestamp NOT NULL
);

CREATE INDEX ON "physician" ("doctorid");

CREATE INDEX ON "physician" ("username");

CREATE INDEX ON "patient" ("patientid");

CREATE INDEX ON "patient" ("username");

CREATE INDEX ON "patientrecords" ("recordid");

CREATE INDEX ON "appointment" ("appointmentid");

ALTER TABLE "patientrecords" ADD FOREIGN KEY ("patientid") REFERENCES "patient" ("patientid") ON DELETE CASCADE;
ALTER TABLE "patientrecords" ADD FOREIGN KEY ("doctorid") REFERENCES "physician" ("doctorid") ON DELETE CASCADE;

ALTER TABLE "appointment" ADD FOREIGN KEY ("doctorid") REFERENCES "physician" ("doctorid") ON DELETE CASCADE;

ALTER TABLE "appointment" ADD FOREIGN KEY ("patientid") REFERENCES "patient" ("patientid") ON DELETE CASCADE;

