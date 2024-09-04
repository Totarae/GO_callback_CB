CREATE SCHEMA IF NOT EXISTS storage_schema;

CREATE TABLE storage_schema.objects (
	id integer NOT NULL,
	last_seen timestamp default current_timestamp,
	CONSTRAINT "pk_object_id" PRIMARY KEY (id)
);

