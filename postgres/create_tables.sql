CREATE TABLE public.users
(
    passport   CHAR(11)     NOT NULL UNIQUE PRIMARY KEY,
    surname    VARCHAR(30)  NOT NULL,
    name       VARCHAR(30)  NOT NULL,
    patronymic VARCHAR(30),
    address    VARCHAR(250) NOT NULL
);

INSERT INTO public.users (passport, surname, name, patronymic, address)
VALUES ('1234 567890', 'Ivanov', 'Ivan', 'Ivanovich', 'Moscow, Lenina 5 apt. 1'),
       ('1534 157798', 'Sherlock', 'Holmes', '', 'London, 221B Backer Street'),
       ('2626 125616', 'Jack', 'Reacher', '', 'Undefined'),
       ('9876 829769', 'Pavlov', 'Konstantin', 'Andreevich', 'Nizhniy Novgorod, Stroiteley 12'),
       ('3265 290523', 'Verstappen', 'Max', '', 'Monaco, 76 Crosby street')
;

CREATE TABLE public.tasks
(
    id            SERIAL PRIMARY KEY,
    user_passport CHAR(11)     NOT NULL,
    task          VARCHAR(100) NOT NULL,
    started_at    TIMESTAMP    NOT NULL,
    finished_at   TIMESTAMP
);

-- INSERT INTO public.tasks (user_id, task, started_at, finished_at)
-- VALUES (1, 'дойка коровы', (now() - INTERVAL '1155 seconds'), now()),
--        (1, 'прополка грядкок', (now() - INTERVAL '10532 seconds'), now()),
--        (1, 'чистка амбара', (now() - INTERVAL '5116 seconds'), now()),
--        (5, 'RB20 tests', (now() - INTERVAL '7626 seconds'), now()),
--        (5, 'sprint qualifying', (now() - INTERVAL '3161 seconds'), now()),
--        (5, 'qualifying', (now() - INTERVAL '4915 seconds'), now()),
--        (5, 'Austrian Grand Prix', (now() - INTERVAL '5166 seconds'), now()),
--        (3, 'murder investigation', (now() - INTERVAL '7235 seconds'), now()),
--        (3, 'fighting in the bar', (now() - INTERVAL '1800 seconds'), now()),
--        (2, 'arguing with Watson', (now() - INTERVAL '4531 seconds'), now()),
--        (2, 'investigating', (now() - INTERVAL '5634 seconds'), now()),
--        (4, 'groceries', (now() - INTERVAL '3235 seconds'), now()),
--        (4, 'plumbing', (now() - INTERVAL '2662 seconds'), now()),
--        (4, 'fixing the porch', (now() - INTERVAL '5432 seconds'), now())
-- ;
