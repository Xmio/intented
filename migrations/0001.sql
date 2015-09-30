CREATE TABLE leads
(
  id bigint NOT NULL,
  mail character varying(100),
  "hashCode" character varying(32),
  invited character varying(32),
  CONSTRAINT leads_pkey PRIMARY KEY (id)
);

CREATE SEQUENCE seq_leads
  INCREMENT 1
  MINVALUE 1
  MAXVALUE 9223372036854775807
  START 1
  CACHE 1;
