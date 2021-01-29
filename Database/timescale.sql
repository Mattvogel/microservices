CREATE database scales;

\connect "scales";

CREATE EXTENSION IF NOT EXISTS timescaledb;

DROP TABLE IF EXISTS "conditions";
CREATE TABLE "public"."conditions" (
    "time" timestamptz NOT NULL,
    "device" text NOT NULL,
    "temperature" double precision,
    "humidity" double precision,
    CONSTRAINT "conditions_time_location_idx" UNIQUE ("time", "device")
) WITH (oids = false);


DROP TABLE IF EXISTS "devices";
DROP SEQUENCE IF EXISTS device_id_seq;
CREATE SEQUENCE device_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 START 3 CACHE 1;

CREATE TABLE "public"."devices" (
    "id" integer DEFAULT nextval('device_id_seq') NOT NULL,
    "owner" character varying NOT NULL,
    "deviceid" character varying NOT NULL,
    "name" character varying,
    "updated_at" integer,
    "created_at" integer,
    CONSTRAINT "devices_id" PRIMARY KEY ("id")
) WITH (oids = false);


DROP TABLE IF EXISTS "user";
DROP SEQUENCE IF EXISTS user_id_seq;
CREATE SEQUENCE user_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 START 1 CACHE 1;

CREATE TABLE "public"."user" (
    "id" integer DEFAULT nextval('user_id_seq') NOT NULL,
    "email" character varying,
    "password" character varying,
    "name" character varying,
    "updated_at" integer,
    "created_at" integer,
    CONSTRAINT "user_id" PRIMARY KEY ("id")
) WITH (oids = false);

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';

CREATE FUNCTION created_at_column() RETURNS trigger
    LANGUAGE plpgsql
    AS $$

    BEGIN
    	NEW.updated_at = EXTRACT(EPOCH FROM NOW());
    	NEW.created_at = EXTRACT(EPOCH FROM NOW());
        RETURN NEW;
    END;
    $$;

CREATE FUNCTION updated_at_column() RETURNS trigger
    LANGUAGE plpgsql
    AS $$

    BEGIN
    	NEW.updated_at = EXTRACT(EPOCH FROM NOW());
        RETURN NEW;
    END;
    $$;

DELIMITER ;;

CREATE TRIGGER "create_user_created_at" BEFORE INSERT ON "public"."user" FOR EACH ROW EXECUTE FUNCTION created_at_column();;

CREATE TRIGGER "update_user_updated_at" BEFORE INSERT ON "public"."user" FOR EACH ROW EXECUTE FUNCTION updated_at_column();;

CREATE TRIGGER "create_user_created_at" BEFORE INSERT ON "public"."devices" FOR EACH ROW EXECUTE FUNCTION created_at_column();;

CREATE TRIGGER "update_user_updated_at" BEFORE INSERT ON "public"."devices" FOR EACH ROW EXECUTE FUNCTION updated_at_column();;
DELIMITER ;
