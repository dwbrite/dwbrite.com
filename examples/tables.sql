-- Drop all tables. This should only be used in the development environment.
drop table posts cascade;
drop table projects cascade;
drop table pages cascade;

-- Example table creation (automated with api).
CREATE TABLE IF NOT EXISTS posts (
  uid serial PRIMARY KEY,
  optional_summary varchar(192),
  title varchar(128) UNIQUE NOT NULL,
  body text NOT NULL,
  post_date date DEFAULT CURRENT_DATE NOT NULL
) WITH (OIDS=FALSE);

-- Example summary function for posts (automated with api).
CREATE OR REPLACE FUNCTION summary(rec posts)
  RETURNS varchar(192)
LANGUAGE SQL
AS
$$
SELECT
       CASE WHEN $1.optional_summary IS NULL
                 THEN
           CASE WHEN length($1.body) > 192
                     THEN
               $1.body::varchar(191) || '…'
                ELSE
               $1.body::varchar(192)
               END
            ELSE
           $1.optional_summary
           END
$$;

-- Example manual insertion into `pages`. Pretty simple, really.
INSERT INTO pages(title, body) VALUES('Home',
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