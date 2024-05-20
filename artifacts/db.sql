CREATE DATABASE todolist;

\c todolist;

CREATE TABLE "User" (
    "UserId" UUID DEFAULT gen_random_uuid(),
    "FirstName" VARCHAR,
    "LastName" VARCHAR,
    "Email" VARCHAR,
    "Login" VARCHAR,
    PRIMARY KEY ("UserId")
);

CREATE TABLE "ToDoList" (
    "ToDoId" UUID DEFAULT gen_random_uuid(),
    "UserId" UUID,
    "Description" VARCHAR,
    "DueDate" TIMESTAMP,
    "Checked" BOOLEAN,
    PRIMARY KEY ("ToDoId"),
    CONSTRAINT fk_user FOREIGN KEY("UserId") REFERENCES "User"("UserId")
);
