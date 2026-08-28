CREATE TABLE "activities" (
	"id" serial PRIMARY KEY,
	"name" text NOT NULL,
	"is_done" boolean DEFAULT false NOT NULL
);
