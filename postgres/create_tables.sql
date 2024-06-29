CREATE TABLE public.users
(
    id              SERIAL PRIMARY KEY,
    passport_series INTEGER,
    passport_number INTEGER,
    surname         VARCHAR(30) NOT NULL,
    name            VARCHAR(30) NOT NULL,
    patronymic      VARCHAR(30),
    address         VARCHAR(250)
);

INSERT INTO public.users (passport_series, passport_number, surname, name, patronymic, address) VALUES
    (1234, 567890, 'Ivanov', 'Ivan', 'Ivanovich', 'Moscow, Lenina 20 apt. 321'),
    (1534, 157798, 'Sherlock', 'Holmes', NULL, 'London, 221B Backer Street'),
    (2626, 125616, 'Jack', 'Reacher', NULL, NULL),
    (9876, 829769, 'Pavlov', 'Konstantin', 'Andreevich', 'Nizhniy Novgorod, Stroiteley 12'),
    (3265, 290523, 'Verstappen', 'Max', NULL, 'Monaco, 76 Crosby street')
;
