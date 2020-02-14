CREATE TABLE likes (
   like_id serial PRIMARY KEY,
   created_by INTEGER REFERENCES users (user_id),
   post_liked INTEGER REFERENCES posts (post_id),
   UNIQUE(created_by, post_liked)
);
