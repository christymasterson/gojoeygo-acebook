CREATE TABLE posts (
   post_id serial PRIMARY KEY,
   created_by INTEGER REFERENCES users (user_id);
   content VARCHAR (500) NOT NULL,
   created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
