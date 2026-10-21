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

INSERT INTO generations (id, name, region) VALUES
 (1, 'generation-i', 'kanto'),
 (2, 'generation-ii', 'johto'),
 (3, 'generation-iii', 'hoenn'),
 (4, 'generation-iv', 'sinnoh'),
 (5, 'generation-v', 'unova'),
 (6, 'generation-vi', 'kalos'),
 (7, 'generation-vii', 'alola'),
 (8, 'generation-viii', 'galar'),
 (9, 'generation-ix', 'paldea')
ON CONFLICT (id) DO UPDATE 
SET name = EXCLUDED.name, region = EXCLUDED.region;

INSERT INTO types (name) VALUES
 ('normal'),('fire'),('water'),('electric'),('grass'),('ice'),
 ('fighting'),('poison'),('ground'),('flying'),('psychic'),('bug'),
 ('rock'),('ghost'),('dragon'),('dark'),('steel'),('fairy')
ON CONFLICT (name) DO NOTHING;

INSERT INTO pokemon (id, name, height, weight, base_exp, hp, attack, defense, sp_atk, sp_def, speed, image_url, generation_id) VALUES
(1, 'bulbasaur', 7, 69, 64, 45, 49, 49, 65, 65, 45, 'https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/1.png', 1),
(2, 'ivysaur', 10, 130, 142, 60, 62, 63, 80, 80, 60, 'https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/2.png', 1),
(3, 'venusaur', 20, 1000, 263, 80, 82, 83, 100, 100, 80, 'https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/3.png', 1),
(4, 'charmander', 6, 85, 62, 39, 52, 43, 60, 50, 65, 'https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/4.png', 1),
(5, 'charmeleon', 11, 190, 142, 58, 64, 58, 80, 65, 80, 'https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/5.png', 1),
(6, 'charizard', 17, 905, 267, 78, 84, 78, 109, 85, 100, 'https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/6.png', 1),
(7, 'squirtle', 5, 90, 63, 44, 48, 65, 50, 64, 43, 'https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/7.png', 1),
(8, 'wartortle', 10, 225, 142, 59, 63, 80, 65, 80, 58, 'https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/8.png', 1),
(9, 'blastoise', 16, 855, 265, 79, 83, 100, 85, 105, 78, 'https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/9.png', 1),
(10, 'caterpie', 3, 29, 39, 45, 30, 35, 20, 20, 45, 'https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/10.png', 1),
(11, 'metapod', 7, 99, 72, 50, 20, 55, 25, 25, 30, 'https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/11.png', 1),
(12, 'butterfree', 11, 320, 198, 60, 45, 50, 90, 80, 70, 'https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/12.png', 1),
(13, 'weedle', 3, 32, 39, 40, 35, 30, 20, 20, 50, 'https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/13.png', 1),
(14, 'kakuna', 6, 100, 72, 45, 25, 50, 25, 25, 35, 'https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/14.png', 1),
(15, 'beedrill', 10, 295, 178, 65, 90, 40, 45, 80, 75, 'https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/15.png', 1),
(16, 'pidgey', 3, 18, 50, 40, 45, 40, 35, 35, 56, 'https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/16.png', 1),
(17, 'pidgeotto', 11, 300, 122, 63, 60, 55, 50, 50, 71, 'https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/17.png', 1),
(18, 'pidgeot', 15, 395, 216, 83, 80, 75, 70, 70, 101, 'https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/18.png', 1),
(19, 'rattata', 3, 35, 51, 30, 56, 35, 25, 35, 72, 'https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/19.png', 1),
(20, 'raticate', 7, 185, 145, 55, 81, 60, 50, 70, 97, 'https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/20.png', 1),
(21, 'spearow', 3, 20, 52, 40, 60, 30, 31, 31, 70, 'https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/21.png', 1),
(22, 'fearow', 12, 380, 155, 65, 90, 65, 61, 61, 100, 'https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/22.png', 1),
(23, 'ekans', 20, 69, 58, 35, 60, 44, 40, 54, 55, 'https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/23.png', 1),
(24, 'arbok', 35, 650, 157, 60, 95, 69, 65, 79, 80, 'https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/24.png', 1),
(25, 'pikachu', 4, 60, 112, 35, 55, 40, 50, 50, 90, 'https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/25.png', 1)
ON CONFLICT (id) DO UPDATE SET
  name = EXCLUDED.name,
  height = EXCLUDED.height,
  weight = EXCLUDED.weight,
  base_exp = EXCLUDED.base_exp,
  hp = EXCLUDED.hp,
  attack = EXCLUDED.attack,
  defense = EXCLUDED.defense,
  sp_atk = EXCLUDED.sp_atk,
  sp_def = EXCLUDED.sp_def,
  speed = EXCLUDED.speed,
  image_url = EXCLUDED.image_url,
  generation_id = EXCLUDED.generation_id,
  updated_at = now();

INSERT INTO pokemon_types (pokemon_id, type_id, slot)
SELECT s.pokemon_id, t.id, s.slot
FROM (VALUES
  (1, 'grass', 1), (1, 'poison', 2),
  (2, 'grass', 1), (2, 'poison', 2),
  (3, 'grass', 1), (3, 'poison', 2),
  (4, 'fire', 1),
  (5, 'fire', 1),
  (6, 'fire', 1), (6, 'flying', 2),
  (7, 'water', 1),
  (8, 'water', 1),
  (9, 'water', 1),
  (10, 'bug', 1),
  (11, 'bug', 1),
  (12, 'bug', 1), (12, 'flying', 2),
  (13, 'bug', 1), (13, 'poison', 2),
  (14, 'bug', 1), (14, 'poison', 2),
  (15, 'bug', 1), (15, 'poison', 2),
  (16, 'normal', 1), (16, 'flying', 2),
  (17, 'normal', 1), (17, 'flying', 2),
  (18, 'normal', 1), (18, 'flying', 2),
  (19, 'normal', 1),
  (20, 'normal', 1),
  (21, 'normal', 1), (21, 'flying', 2),
  (22, 'normal', 1), (22, 'flying', 2),
  (23, 'poison', 1),
  (24, 'poison', 1),
  (25, 'electric', 1)
) AS s(pokemon_id, type_name, slot)
JOIN types t ON t.name = s.type_name
ON CONFLICT (pokemon_id, slot) DO UPDATE SET type_id = EXCLUDED.type_id;


INSERT INTO type_effectiveness (attacker_type_id, defender_type_id, multiplier)
SELECT a.id, d.id, 1.0
FROM types a CROSS JOIN types d
ON CONFLICT (attacker_type_id, defender_type_id) DO UPDATE SET multiplier = EXCLUDED.multiplier;

INSERT INTO type_effectiveness (attacker_type_id, defender_type_id, multiplier)
SELECT a.id, d.id, v.mult
FROM (VALUES
  ('normal', 'rock', 0.5), ('normal', 'steel', 0.5), ('normal', 'ghost', 0.0),
  ('fire', 'grass', 2.0), ('fire', 'ice', 2.0), ('fire', 'bug', 2.0), ('fire', 'steel', 2.0),
  ('fire', 'fire', 0.5), ('fire', 'water', 0.5), ('fire', 'rock', 0.5), ('fire', 'dragon', 0.5),
  ('water', 'fire', 2.0), ('water', 'ground', 2.0), ('water', 'rock', 2.0),
  ('water', 'water', 0.5), ('water', 'grass', 0.5), ('water', 'dragon', 0.5),
  ('electric', 'water', 2.0), ('electric', 'flying', 2.0),
  ('electric', 'electric', 0.5), ('electric', 'grass', 0.5), ('electric', 'dragon', 0.5),
  ('electric', 'ground', 0.0),
  ('grass', 'water', 2.0), ('grass', 'ground', 2.0), ('grass', 'rock', 2.0),
  ('grass', 'fire', 0.5), ('grass', 'grass', 0.5), ('grass', 'poison', 0.5), ('grass', 'flying', 0.5), ('grass', 'bug', 0.5), ('grass', 'dragon', 0.5), ('grass', 'steel', 0.5),
  ('ice', 'grass', 2.0), ('ice', 'ground', 2.0), ('ice', 'flying', 2.0), ('ice', 'dragon', 2.0),
  ('ice', 'fire', 0.5), ('ice', 'water', 0.5), ('ice', 'ice', 0.5), ('ice', 'steel', 0.5),
  ('fighting', 'normal', 2.0), ('fighting', 'ice', 2.0), ('fighting', 'rock', 2.0), ('fighting', 'dark', 2.0), ('fighting', 'steel', 2.0),
  ('fighting', 'poison', 0.5), ('fighting', 'flying', 0.5), ('fighting', 'psychic', 0.5), ('fighting', 'bug', 0.5), ('fighting', 'fairy', 0.5),
  ('fighting', 'ghost', 0.0),
  ('poison', 'grass', 2.0), ('poison', 'fairy', 2.0),
  ('poison', 'poison', 0.5), ('poison', 'ground', 0.5), ('poison', 'rock', 0.5), ('poison', 'ghost', 0.5),
  ('poison', 'steel', 0.0),
  ('ground', 'fire', 2.0), ('ground', 'electric', 2.0), ('ground', 'poison', 2.0), ('ground', 'rock', 2.0), ('ground', 'steel', 2.0),
  ('ground', 'grass', 0.5), ('ground', 'bug', 0.5),
  ('ground', 'flying', 0.0),
  ('flying', 'grass', 2.0), ('flying', 'fighting', 2.0), ('flying', 'bug', 2.0),
  ('flying', 'electric', 0.5), ('flying', 'rock', 0.5), ('flying', 'steel', 0.5),
  ('psychic', 'fighting', 2.0), ('psychic', 'poison', 2.0),
  ('psychic', 'psychic', 0.5), ('psychic', 'steel', 0.5),
  ('psychic', 'dark', 0.0),
  ('bug', 'grass', 2.0), ('bug', 'psychic', 2.0), ('bug', 'dark', 2.0),
  ('bug', 'fire', 0.5), ('bug', 'fighting', 0.5), ('bug', 'poison', 0.5), ('bug', 'flying', 0.5), ('bug', 'ghost', 0.5), ('bug', 'steel', 0.5), ('bug', 'fairy', 0.5),
  ('rock', 'fire', 2.0), ('rock', 'ice', 2.0), ('rock', 'flying', 2.0), ('rock', 'bug', 2.0),
  ('rock', 'fighting', 0.5), ('rock', 'ground', 0.5), ('rock', 'steel', 0.5),
  ('ghost', 'psychic', 2.0), ('ghost', 'ghost', 2.0),
  ('ghost', 'dark', 0.5),
  ('ghost', 'normal', 0.0),
  ('dragon', 'dragon', 2.0),
  ('dragon', 'steel', 0.5),
  ('dragon', 'fairy', 0.0),
  ('dark', 'psychic', 2.0), ('dark', 'ghost', 2.0),
  ('dark', 'fighting', 0.5), ('dark', 'dark', 0.5), ('dark', 'fairy', 0.5),
  ('steel', 'ice', 2.0), ('steel', 'rock', 2.0), ('steel', 'fairy', 2.0),
  ('steel', 'fire', 0.5), ('steel', 'water', 0.5), ('steel', 'electric', 0.5), ('steel', 'steel', 0.5),
  ('fairy', 'fighting', 2.0), ('fairy', 'dragon', 2.0), ('fairy', 'dark', 2.0),
  ('fairy', 'fire', 0.5), ('fairy', 'poison', 0.5), ('fairy', 'steel', 0.5)
) AS v(att, def, mult)
JOIN types a ON a.name = v.att
JOIN types d ON d.name = v.def
ON CONFLICT (attacker_type_id, defender_type_id) DO UPDATE SET multiplier = EXCLUDED.multiplier;
