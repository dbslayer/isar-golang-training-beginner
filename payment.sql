-- public.payment_codes definition

-- Drop table

-- DROP TABLE public.payment_codes;

CREATE TABLE public.payment_codes (
	id uuid NOT NULL,
	payment_code varchar NULL,
	"name" varchar NULL,
	status varchar NULL,
	expiration_date varchar NULL,
	created_date date NULL,
	updated_at date NULL,
	CONSTRAINT payment_codes_pk PRIMARY KEY (id)
);