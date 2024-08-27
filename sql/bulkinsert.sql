-- RESTAURANTS
INSERT INTO restaurants(
	name, dietary_restrictions)
	VALUES ('Raos', '{"Vegetarian-Friendly", "Gluten Free Options"}');
INSERT INTO restaurants(
	name, dietary_restrictions)
	VALUES ('Wild Ginger', '{"Vegetarian-Friendly", "Vegan", "Gluten Free Options"}');
INSERT INTO restaurants(
	name, dietary_restrictions)
	VALUES ('Muscle Maker', '{"Keto"}');

-- TABLES
INSERT INTO public.tables(capacity, restaurant_id) VALUES (2,3);

-- Users
INSERT INTO public.users(
	first_name, last_name, phone_number)
	VALUES ('Jane', 'Flintstone', '+18885555555');
INSERT INTO public.users(
	first_name, last_name, phone_number)
	VALUES ('George', 'Jetson', '+18885551111');

-- RESERVATIONS
INSERT INTO reservations(
	name, party_size, "time", restaurant_id, table_id, user_id)
	VALUES ('Flintstone Family', 2, '2024-08-30T23:00:00Z', 1, 2, 1);
INSERT INTO reservations(
	name, party_size, "time", restaurant_id, table_id, user_id)
	VALUES ('Jetson Family', 2, '2024-08-31T23:00:00Z', 2, 4, 2);