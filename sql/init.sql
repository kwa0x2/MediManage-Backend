DO $$
BEGIN

CREATE TYPE public.user_role_type AS ENUM
    ('Staff', 'Worker');


CREATE TABLE IF NOT EXISTS public."Province"
(
    province_id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
    province_name character varying COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT "Province_pkey" PRIMARY KEY (province_id),
    CONSTRAINT "Province_province_name_key" UNIQUE NULLS NOT DISTINCT (province_name)
    );


CREATE TABLE IF NOT EXISTS public."District"
(
    district_id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
    district_name character varying COLLATE pg_catalog."default" NOT NULL,
    province_name character varying COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT "District_pkey" PRIMARY KEY (district_id),
    CONSTRAINT "District_district_name_key" UNIQUE NULLS NOT DISTINCT (district_name),
    CONSTRAINT "District_district_name_province_name_key" UNIQUE NULLS NOT DISTINCT (district_name, province_name),
    CONSTRAINT "District_province_name_fkey" FOREIGN KEY (province_name)
    REFERENCES public."Province" (province_name) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID
    );


CREATE TABLE IF NOT EXISTS public."Clinic"
(
    clinic_id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
    clinic_name character varying COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT "Clinic_pkey" PRIMARY KEY (clinic_id),
    CONSTRAINT "Clinic_clinic_name_key" UNIQUE NULLS NOT DISTINCT (clinic_name)
    );

CREATE TABLE IF NOT EXISTS public."JobGroup"
(
    jobgroup_id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
    jobgroup_name character varying COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT "JobGroup_pkey" PRIMARY KEY (jobgroup_id),
    CONSTRAINT "JobGroup_jobgroup_name_key" UNIQUE NULLS NOT DISTINCT (jobgroup_name)
    );

CREATE TABLE IF NOT EXISTS public."Employee"
(
    employee_id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
    employee_name character varying COLLATE pg_catalog."default" NOT NULL,
    employee_surname character varying COLLATE pg_catalog."default" NOT NULL,
    employee_identity_number bigint NOT NULL,
    employee_phone_number character varying(15) COLLATE pg_catalog."default" NOT NULL,
    employee_job_group_name character varying COLLATE pg_catalog."default" NOT NULL,
    employee_title_name character varying COLLATE pg_catalog."default" NOT NULL,
    employee_clinic_name character varying COLLATE pg_catalog."default",
    employee_working_days character varying COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT "Staff_pkey" PRIMARY KEY (employee_id),
    CONSTRAINT "Staff_staff_identity_number_staff_phone_number_key" UNIQUE NULLS NOT DISTINCT (employee_identity_number, employee_phone_number)
    );

CREATE TABLE IF NOT EXISTS public."Title"
(
    title_id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
    title_name character varying COLLATE pg_catalog."default" NOT NULL,
    title_jobgroup_name character varying COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT "Title_pkey" PRIMARY KEY (title_id),
    CONSTRAINT "Title_title_name_title_jobgroup_name_key" UNIQUE NULLS NOT DISTINCT (title_name, title_jobgroup_name),
    CONSTRAINT "Title_title_jobgroup_name_fkey" FOREIGN KEY (title_jobgroup_name)
    REFERENCES public."JobGroup" (jobgroup_name) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID
    );

CREATE TABLE IF NOT EXISTS public."User"
(
    user_id uuid NOT NULL DEFAULT gen_random_uuid(),
    user_name character varying COLLATE pg_catalog."default" NOT NULL,
    user_surname character varying COLLATE pg_catalog."default" NOT NULL,
    user_identity_number character varying(11) COLLATE pg_catalog."default" NOT NULL,
    user_email character varying COLLATE pg_catalog."default" NOT NULL,
    user_phone character varying(15) COLLATE pg_catalog."default" NOT NULL,
    user_password character varying COLLATE pg_catalog."default" NOT NULL,
    user_role user_role_type NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    CONSTRAINT "User_pkey" PRIMARY KEY (user_id),
    CONSTRAINT "User_user_email_key" UNIQUE NULLS NOT DISTINCT (user_email),
    CONSTRAINT "User_user_identity_number_key" UNIQUE NULLS NOT DISTINCT (user_identity_number),
    CONSTRAINT "User_user_phone_key" UNIQUE NULLS NOT DISTINCT (user_phone)
    );

CREATE TABLE IF NOT EXISTS public."Hospital"
(
    hospital_id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
    hospital_name character varying COLLATE pg_catalog."default" NOT NULL,
    hospital_tax_identity_number character varying(10) COLLATE pg_catalog."default" NOT NULL,
    hospital_email character varying COLLATE pg_catalog."default" NOT NULL,
    hospital_phone_number character varying(15) COLLATE pg_catalog."default" NOT NULL,
    hospital_province_name character varying COLLATE pg_catalog."default" NOT NULL,
    hospital_district_name character varying COLLATE pg_catalog."default" NOT NULL,
    hospital_address text COLLATE pg_catalog."default" NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    CONSTRAINT "Hospital_pkey" PRIMARY KEY (hospital_id),
    CONSTRAINT "Hospital_hospital_email_key" UNIQUE NULLS NOT DISTINCT (hospital_email),
    CONSTRAINT "Hospital_hospital_phone_number_key" UNIQUE NULLS NOT DISTINCT (hospital_phone_number),
    CONSTRAINT "Hospital_hospital_tax_identity_number_key" UNIQUE NULLS NOT DISTINCT (hospital_tax_identity_number),
    CONSTRAINT "Hospital_hospital_district_fkey" FOREIGN KEY (hospital_district_name)
    REFERENCES public."District" (district_name) MATCH SIMPLE
                         ON UPDATE NO ACTION
                         ON DELETE NO ACTION
    NOT VALID,
    CONSTRAINT "Hospital_hospital_province_fkey" FOREIGN KEY (hospital_province_name)
    REFERENCES public."Province" (province_name) MATCH SIMPLE
                         ON UPDATE NO ACTION
                         ON DELETE NO ACTION
    NOT VALID
    );

INSERT INTO "Province" (province_name) VALUES
                       ('İzmir'),
                       ('İstanbul'),
                       ('Bursa'),
                       ('Ankara');

INSERT INTO "District" (district_name, province_name) VALUES
                      ('Çankaya', 'Ankara'),
                      ('Altındağ', 'Ankara'),
                      ('Etimesgut', 'Ankara'),
                      ('Keçiören', 'Ankara'),
                      ('Nilüfer', 'Bursa'),
                      ('Mudanya', 'Bursa'),
                      ('Osmangazi', 'Bursa'),
                      ('Yıldırım', 'Bursa'),
                      ('Avcılar', 'İstanbul'),
                      ('Büyükçekmece', 'İstanbul'),
                      ('Beylikdüzü', 'İstanbul'),
                      ('Bornova', 'İzmir'),
                      ('Buca', 'İzmir'),
                      ('Çeşme', 'İzmir'),
                      ('Menemen', 'İzmir'),
                      ('Ataşehir', 'İstanbul');

INSERT INTO "JobGroup" (jobgroup_name) VALUES
                       ('Hizmet Personeli'),
                       ('İdari Personel'),
                       ('Doktor');

INSERT INTO "Title" (title_name,title_jobgroup_name) VALUES
                     ('Uzman', 'Doktor'),
                     ('Asistan', 'Doktor'),
                     ('Temizlikçi', 'Hizmet Personeli'),
                     ('Danışman', 'Hizmet Personeli'),
                     ('Başhekim', 'İdari Personel');

EXCEPTION
    WHEN OTHERS THEN
        ROLLBACK;
        RAISE;
END $$;