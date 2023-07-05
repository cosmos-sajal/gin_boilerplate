
-- +migrate Up
CREATE INDEX idx_mobile_number ON public.users USING btree (mobile_number);

-- +migrate Down
DROP INDEX idx_mobile_number;

