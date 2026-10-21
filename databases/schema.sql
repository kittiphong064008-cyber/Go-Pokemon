DROP TABLE IF EXISTS pokemons CASCADE;

CREATE TABLE IF NOT EXISTS generations (
    id     SMALLINT PRIMARY KEY,
    name   VARCHAR(30) UNIQUE NOT NULL,
    region VARCHAR(30)
);

CREATE TABLE IF NOT EXISTS types (
    id   SERIAL PRIMARY KEY,
    name VARCHAR(30) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS type_effectiveness (
    attacker_type_id INT NOT NULL REFERENCES types(id) ON DELETE CASCADE,
    defender_type_id INT NOT NULL REFERENCES types(id) ON DELETE CASCADE,
    multiplier       NUMERIC(2,1) NOT NULL CHECK (multiplier IN (0, 0.5, 1, 2)),
    PRIMARY KEY (attacker_type_id, defender_type_id)
);

CREATE TABLE IF NOT EXISTS abilities (
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(50) UNIQUE NOT NULL,
    description TEXT
);

CREATE TABLE IF NOT EXISTS moves (
    id       SERIAL PRIMARY KEY,
    name     VARCHAR(50) UNIQUE NOT NULL,
    type_id  INT NOT NULL REFERENCES types(id),
    category VARCHAR(10) NOT NULL CHECK (category IN ('physical', 'special', 'status')),
    power    SMALLINT,
    accuracy SMALLINT,
    pp       SMALLINT
);

CREATE TABLE IF NOT EXISTS pokemon (
    id            INT PRIMARY KEY,
    name          VARCHAR(50) UNIQUE NOT NULL,
    height        INT NOT NULL,
    weight        INT NOT NULL,
    base_exp      INT,
    hp            SMALLINT NOT NULL,
    attack        SMALLINT NOT NULL,
    defense       SMALLINT NOT NULL,
    sp_atk        SMALLINT NOT NULL,
    sp_def        SMALLINT NOT NULL,
    speed         SMALLINT NOT NULL,
    image_url     TEXT,
    generation_id SMALLINT REFERENCES generations(id),
    created_at    TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS pokemon_types (
    pokemon_id INT NOT NULL REFERENCES pokemon(id) ON DELETE CASCADE,
    type_id    INT NOT NULL REFERENCES types(id),
    slot       SMALLINT NOT NULL CHECK (slot IN (1, 2)),
    PRIMARY KEY (pokemon_id, type_id),
    UNIQUE (pokemon_id, slot)
);

CREATE TABLE IF NOT EXISTS pokemon_abilities (
    pokemon_id INT NOT NULL REFERENCES pokemon(id) ON DELETE CASCADE,
    ability_id INT NOT NULL REFERENCES abilities(id),
    is_hidden  BOOLEAN NOT NULL DEFAULT false,
    PRIMARY KEY (pokemon_id, ability_id)
);

CREATE TABLE IF NOT EXISTS pokemon_moves (
    pokemon_id  INT NOT NULL REFERENCES pokemon(id) ON DELETE CASCADE,
    move_id     INT NOT NULL REFERENCES moves(id),
    learn_level SMALLINT,
    PRIMARY KEY (pokemon_id, move_id)
);

CREATE TABLE IF NOT EXISTS evolutions (
    from_pokemon_id INT NOT NULL REFERENCES pokemon(id) ON DELETE CASCADE,
    to_pokemon_id   INT NOT NULL REFERENCES pokemon(id) ON DELETE CASCADE,
    min_level       SMALLINT,
    trigger         VARCHAR(30),
    PRIMARY KEY (from_pokemon_id, to_pokemon_id),
    CHECK (from_pokemon_id <> to_pokemon_id)
);

CREATE INDEX IF NOT EXISTS idx_pokemon_name           ON pokemon (name);
CREATE INDEX IF NOT EXISTS idx_pokemon_generation     ON pokemon (generation_id);
CREATE INDEX IF NOT EXISTS idx_pokemon_types_type     ON pokemon_types (type_id);
CREATE INDEX IF NOT EXISTS idx_pokemon_abilities_abl  ON pokemon_abilities (ability_id);
CREATE INDEX IF NOT EXISTS idx_pokemon_moves_move     ON pokemon_moves (move_id);
CREATE INDEX IF NOT EXISTS idx_moves_type             ON moves (type_id);
CREATE INDEX IF NOT EXISTS idx_evolutions_to          ON evolutions (to_pokemon_id);

INSERT INTO generations (id, name, region) 
VALUES (1, 'generation-i', 'kanto')
ON CONFLICT (id) DO NOTHING;

INSERT INTO types (name) VALUES
 ('normal'),('fire'),('water'),('electric'),('grass'),('ice'),
 ('fighting'),('poison'),('ground'),('flying'),('psychic'),('bug'),
 ('rock'),('ghost'),('dragon'),('dark'),('steel'),('fairy')
ON CONFLICT (name) DO NOTHING;

INSERT INTO pokemon (id, name, height, weight, base_exp, hp, attack, defense, sp_atk, sp_def, speed, image_url, generation_id)
VALUES (25, 'pikachu', 4, 60, 112, 35, 55, 40, 50, 50, 90, 'https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/25.png', 1)
ON CONFLICT (id) DO NOTHING;

INSERT INTO pokemon_types (pokemon_id, type_id, slot)
SELECT 25, id, 1 FROM types WHERE name = 'electric'
ON CONFLICT (pokemon_id, slot) DO NOTHING;
