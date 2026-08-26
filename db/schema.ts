import { boolean, pgTable, serial, text } from "drizzle-orm/pg-core";

export const activities = pgTable("activities", {
  id: serial().primaryKey(),
  name: text().notNull(),
  isDone: boolean("is_done").notNull().default(false),
});
