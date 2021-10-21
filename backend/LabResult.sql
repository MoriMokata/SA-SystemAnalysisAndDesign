BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "medical_teches" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"name"	text,
	"email"	text,
	"pass"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "medical_records" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"hospital_number"	text,
	"personal_id"	text,
	"patient_name"	text,
	"patient_age"	integer,
	"patient_dob"	datetime,
	"patient_tel"	text,
	"register_date"	datetime,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "lab_types" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"name"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "lab_rooms" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"name"	text,
	"building"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "lab_results" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"medical_tech_id"	integer,
	"medical_record_id"	integer,
	"lab_type_id"	integer,
	"lab_result"	text,
	"lab_detail"	text,
	"lab_room_id"	integer,
	"added_time"	datetime,
	PRIMARY KEY("id"),
	CONSTRAINT "fk_medical_teches_lab_results" FOREIGN KEY("medical_tech_id") REFERENCES "medical_teches"("id"),
	CONSTRAINT "fk_medical_records_lab_results" FOREIGN KEY("medical_record_id") REFERENCES "medical_records"("id"),
	CONSTRAINT "fk_lab_rooms_lab_results" FOREIGN KEY("lab_room_id") REFERENCES "lab_rooms"("id"),
	CONSTRAINT "fk_lab_types_lab_results" FOREIGN KEY("lab_type_id") REFERENCES "lab_types"("id")
);
CREATE UNIQUE INDEX IF NOT EXISTS "idx_medical_teches_email" ON "medical_teches" (
	"email"
);
CREATE INDEX IF NOT EXISTS "idx_medical_teches_deleted_at" ON "medical_teches" (
	"deleted_at"
);
CREATE UNIQUE INDEX IF NOT EXISTS "idx_medical_records_personal_id" ON "medical_records" (
	"personal_id"
);
CREATE INDEX IF NOT EXISTS "idx_medical_records_deleted_at" ON "medical_records" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_lab_types_deleted_at" ON "lab_types" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_lab_rooms_deleted_at" ON "lab_rooms" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_lab_results_deleted_at" ON "lab_results" (
	"deleted_at"
);
COMMIT;
