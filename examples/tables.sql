-- Drop all tables. This should only be used in the development environment.
drop table posts;
drop table projects;
drop table pages;

-- Create all tables.
CREATE TABLE posts (
  uid serial PRIMARY KEY,
  title varchar(128) NOT NULL,
  body text NOT NULL,
  post_date date DEFAULT CURRENT_DATE
) WITH (OIDS=FALSE);

CREATE TABLE projects (
  title varchar(128) PRIMARY KEY,
  summary varchar(192) NOT NULL,
  body text NOT NULL
) WITH (OIDS=FALSE);

CREATE TABLE pages (
  title varchar(128) PRIMARY KEY,
  body text NOT NULL
) WITH (OIDS=FALSE);

-- Example insertion into `posts`.
-- INSERT INTO posts(title, body) VALUES('o', 'hej');

-- Example insertions into `pages`.
INSERT INTO pages(title, body) VALUES('About Me',
                                      '<p/>Hello! My name is Devin Brite. ' ||
                                      'As a college student I''m involved in web development and programming. ' ||
                                      'As a hobbyist, I make games. ' ||
                                      'I also fight for personal freedoms - ' ||
                                      'be that in the name of free software, social rights, privacy, or the open Internet. ' ||
                                      'Despite being an advocate for privacy, I consider myself a very public person. ' ||
                                      'Speaking of, you can find all my social media accounts at ' ||
                                      '<a href="https://keybase.io/hd">https://keybase.io/hd.</a>. Yes, all.' ||
                                      '' ||
                                      '<p/>Here''s a brief history of me:' ||
                                      '' ||
                                      '<p/>I grew up in a town called Millville, Massachusetts. ' ||
                                      'I moved to Grafton in my pre-teen years, and I''ve lived here since. ' ||
                                      'I dropped out of high school, took a year off, ' ||
                                      'and ended up with a diploma at the same time as my peers. <i>Eat that, world.</i> ' ||
                                      'Now I''m getting an Associate''s in Science, majoring in Computer Information Systems.' ||
                                      '' ||
                                      '<p/>Now let''s talk technology. ' ||
                                      'I''m most experienced with Java, but my favorite programming language is Rust. ' ||
                                      'My least favorite is Visual Basic. I <i>despise</i> the Microsoft ecosystem, ' ||
                                      'and I will tell you this every time I have to use it. ' ||
                                      '<a href="https://docs.google.com/spreadsheets/d/1R9KzpWbyOp9GM-laW3as4nkQBD2jWWaAT8XgrpFzJN0/edit?usp=sharing">I <i>really</i> like spreadsheets.</a>');


INSERT INTO pages(title, body) VALUES('401', 'Get out. Get out right now. You''re not welcome here. Unless you are. In which case... You may want to update your credentials.');
INSERT INTO pages(title, body) VALUES('403', 'Quit snooping around!');
INSERT INTO pages(title, body) VALUES('404', 'Page not found :(');
INSERT INTO pages(title, body) VALUES('503', 'Either you''re overloading my server or it''s down for maintenance. Chill.');