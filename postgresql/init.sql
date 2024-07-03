CREATE USER postgres WITH SUPERUSER PASSWORD 'password';

CREATE TABLE public.users
(
    id          SERIAL PRIMARY KEY,
    email       character varying(255),
    first_name  character varying(255),
    last_name   character varying(255),
    password    character varying(60),
    user_active integer DEFAULT 0,
    created_at  timestamp without time zone,
    updated_at  timestamp without time zone
);

CREATE TABLE public.auth_refresh_tokens
(
    id         SERIAL PRIMARY KEY,
    token      VARCHAR UNIQUE NOT NULL,
    user_id    INT            NOT NULL,
    active     BOOLEAN        NOT NULL,
    used_count SMALLINT       NOT NULL,
    created_at TIMESTAMP      NOT NULL,
    expires_at TIMESTAMP      NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE public.articles
(
    id         SERIAL PRIMARY KEY,
    image      VARCHAR(255),
    title      VARCHAR(255) NOT NULL,
    text       TEXT         NOT NULL,
    author_id  INT          NOT NULL,
    created_at TIMESTAMP    NOT NULL,
    updated_at TIMESTAMP    NOT NULL,
    FOREIGN KEY (author_id) REFERENCES users (id)
);

ALTER TABLE public.users OWNER TO postgres;
ALTER TABLE public.auth_refresh_tokens OWNER TO postgres;
ALTER TABLE public.articles OWNER TO postgres;


INSERT INTO "public"."users"("email", "first_name", "last_name", "password", "user_active", "created_at", "updated_at")
VALUES ('admin@example.com', 'Admin', 'User', '$2a$12$1zGLuYDDNvATh4RA4avbKuheAMpb1svexSzrQm7up.bnpwQHs0jNe', 1,
        '2022-03-14 00:00:00', '2022-03-14 00:00:00');
INSERT INTO public.articles (id, image, title, text, author_id, created_at, updated_at) VALUES (1, 'https://miro.medium.com/v2/resize:fit:1000/format:webp/1*HOVbioEdQ1rcfoMcPorIkw@2x.jpeg', 'The Impact of Travel on Creativity', e'As writers, our craft is deeply influenced by our experiences. Among these, travel stands out as a powerful catalyst for creativity. The unfamiliar sights, sounds, and sensations we encounter while exploring new places can profoundly reshape our narratives and expand our creative horizons.

Cultural Immersion and Perspective Shift
Immersing ourselves in different cultures challenges our preconceptions and broadens our worldview. This shift in perspective often translates into more nuanced and empathetic storytelling. Writers like Ernest Hemingway and Elizabeth Gilbert have famously drawn from their travels to create rich, culturally diverse narratives.

2. Sensory Stimulation and Descriptive Writing

New environments bombard us with fresh sensory inputs. The scent of spices in a Moroccan market, the cacophony of a Tokyo street, or the texture of sand in the Sahara — these vivid experiences enhance our ability to craft detailed, evocative descriptions that bring our stories to life.

3. Overcoming Creative Blocks

Stepping out of our routine and comfort zone can be an effective way to overcome writer’s block. Travel disrupts our patterns, forcing our brains to form new neural connections. This mental shake-up often leads to breakthrough ideas and novel plot twists.

4. Character Inspiration

Every journey introduces us to a cast of unique individuals. These encounters can inspire compelling fictional characters or add depth to existing ones. Travel allows us to observe and understand diverse human behaviors, enriching our character development skills.', 1, '2024-07-03 12:17:07.000000', '2024-07-03 12:17:08.000000');
INSERT INTO public.articles (id, image, title, text, author_id, created_at, updated_at) VALUES (2, 'https://miro.medium.com/v2/resize:fit:4800/format:webp/0*Q00r9TP_wzVLbADc', 'How I Make Money On Medium Without MPP', e'I published my first article on Medium in May 2022.

After hitting 100 followers I realized my country was delisted from MPP. It will come back soon but that’s another story.

For now, anyone can read my stories for free. And I know many people are in a similar condition right now.

Let me ask you something.

How many social media platforms are currently available?

YouTube. Twitter. Instagram. Pinterest. TikTok. Medium. Reddit. To name a few.

How many of these platforms have a Partner Program available in every country?

Only ONE. YOUTUBE.

I thought of becoming a YouTuber when I wasn’t able to join MPP.

But then I remember something.

There is not only one way to make money through social media platforms. Not all success stories out there mentioned only Partner’s program as their main income source.

', 1, '2024-07-03 13:00:21.000000', '2024-07-03 13:00:24.000000');



