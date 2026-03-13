#Name 
Thet Htet Oo

#Link to github repo
https://github.com/Thet-Htet-Oo/gossip.git

#Link to live app
https://gossip-frontend.onrender.com
https://gossip-backend-1nwv.onrender.com

#Live development
You can open the link https://gossip-backend-1nwv.onrender.com first in the browser. If you see 404 , then you are ready your backend
and you can easily try my app with the link https://gossip-frontend.onrender.com .

The backend and frontend are deployed using Render's free hosting tier.
Due to inactivity, the backend service may enter a sleep state. When accessed again, the server may take a short time to restart.
You may open the backend URL first to wake the server before accessing the frontend.

If a Cloudflare error appears:
1. Refresh the backend URL 
2. Please wait about 30-60 seconds
3. Then open the frontend

You can also try localhost since I provided below. 
///////////////////////////////////////////////////////////////////////////////////////////////////////////////
Use of AI

- Since I am not friendly to GO and TypeScript, I have to ask the AI for step by step guideline to learn first.
- Then, I watched YouTube videos and wrote my codes.
- When I didn't know how to write, I used AI to teach me.
- When I faced with errors, I tried to find the solution myself first. If it was a very complex and I couldn't solve it anymore
  then, I asked AI why this error happened and what was the couse of this errors . I read the explanation from AI and solved the errors.
- I also used AI to review my code and give me better suggesstions.
  
///////////////////////////////////////////////////////////////////////////////////////////////////////////////
Features
This application allows users to:
- Create topics
- Create posts under topics
- Comment on posts
- Reply to comments
- Like posts
- Update and delete topics
- Update and delete posts
- Update and delete comments

//////////////////////////////////////////////////////////////////////////////////////////////////////////////
#Local Development
This is optional only if you want to test this on localhost!!!!
If you want to run with localhost use this repo link : https://github.com/Thet-Htet-Oo/gossip-local.git
1. Install Go from https://golang.org/dl/
2. Then check go version in command prompt
3. If you already have Go , run " go run main.go " in command prompt from backend directory

4. For frontend , cd frontend
5. npm install
6. npm install @mui/material @emotion/react @emotion/styled @mui/icons-material
7. npm start 

For database, in command prompt 
1. psql -U postgres
2. Then enter your password
3. CREATE DATABASE gossip;
4. \c gossip
5. CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
6. CREATE TABLE topics (
    id SERIAL PRIMARY KEY,       
    title VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
7. CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    topic_id INT REFERENCES topics(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    content TEXT,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
8. CREATE TABLE comments (
    id SERIAL PRIMARY KEY,
    post_id INT REFERENCES posts(id) ON DELETE CASCADE,
    parent_comment_id INT REFERENCES comments(id) ON DELETE CASCADE,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
9. CREATE TABLE post_likes (
    id SERIAL PRIMARY KEY,
    post_id INT REFERENCES posts(id) ON DELETE CASCADE,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    UNIQUE(post_id, user_id)  
);
10. Important ** Try to adjust the .env file as you need.
11. Then run " go run main.go " 
12. You will see Database connected successfully and Server running on http://localhost:8000
13. Then run " npm start "
14. This will lead you to http://localhost:3000 and you can start using.  
