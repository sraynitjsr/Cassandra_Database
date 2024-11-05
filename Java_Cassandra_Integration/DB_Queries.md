-- 1. List All Users
SELECT * FROM my_keyspace.users;

-- 2. Get Specific User by ID
SELECT * FROM my_keyspace.users WHERE id = 10;

-- 3. Create a New User
INSERT INTO my_keyspace.users (id, first_name, last_name, email) VALUES (uuid(), 'Aarav', 'Sharma', 'aarav.sharma@example.com');

-- 4. Update User by ID
UPDATE my_keyspace.users SET first_name = 'Aarav', last_name = 'Sharma', email = 'aarav.sharma@newdomain.com' WHERE id = 25;

-- 5. Delete User by ID 
DELETE FROM my_keyspace.users WHERE id = 35;

-- 6. Check If User Exists
SELECT * FROM my_keyspace.users WHERE id = 420;

-- 7. Count the Number of Users
SELECT count(*) FROM my_keyspace.users;

-- 8. Drop the User Table
DROP TABLE my_keyspace.users;
