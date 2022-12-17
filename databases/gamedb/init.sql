CREATE TABLE alliance(
    alliance_id INT NOT NULL,
    alliance_name VARCHAR(50),
    alliance_description TEXT,

    PRIMARY KEY(alliance_id)
);

CREATE TABLE player(
    player_id INT NOT NULL,
    nickname VARCHAR(50),
    title VARCHAR(50),
    points INT,
    -- Each bit will indicate boolean for completed research (BIGINT = 64 bits).
    research BIGINT,
    alliance_id INT REFERENCES alliance (alliance_id),
    
    PRIMARY KEY(player_id)
);

CREATE TYPE alliance_relation AS ENUM('neutral', 'war', 'allied', 'peace-treaty');

CREATE TABLE alliance_relations(
    a_alliance_id INT NOT NULL REFERENCES alliance (alliance_id),
    b_alliance_id INT NOT NULL REFERENCES alliance (alliance_id),
    relation_type alliance_relation,

    PRIMARY KEY(a_alliance_id, b_alliance_id)
);

CREATE TABLE city(
    city_id INT NOT NULL,
    player_id INT NOT NULL REFERENCES player (player_id),
    -- Location ID is will use 2 Bytes for X and 2 Bytes for Y.
    location_id INT,
    city_name VARCHAR(50),
    r_food_amount INT,
    r_stick_amount INT,
    r_rock_amount INT,
    r_clay_amount INT,
    r_coin_amount INT,
    u_swordsman_amount INT,
    u_stickoplite_amount INT,
    u_pikeman_amount INT,
    u_archer_amount INT,
    u_crossbro_amount INT,
    u_slinger_amount INT,
    u_centaur_amount INT,
    u_spidosaur_amount INT,
    u_stickofants_amount INT,
    u_trebuchet_amount INT,
    u_ram_amount INT,
    u_culverin_amount INT,
    u_ciroustick_amount INT,
    u_caravan_amount INT,
    u_colonizer_amount INT,
    u_fishboat_amount INT,
    u_scientist_amount INT,
    b_city_hall_level INT,
    b_statue_level INT,
    b_barracks_level INT,
    b_wall_level INT,
    b_quarry_level INT,
    b_lumbermill_level INT,
    b_clay_pit_level INT,
    b_farm_level INT,
    b_warehouse_level INT,
    b_market_level INT,
    b_fishing_spot_level INT,
    b_bank_level INT,
    b_embassy_level INT,
    b_library_level INT,
    b_research_lab_level INT,
    b_hospital_level INT,
    b_harbor_level INT,
    b_airport_level INT,
    b_mutant_lab_level INT,
    b_temple_level INT,
    income_tax FLOAT,
    update_timestamp TIMESTAMP,

    PRIMARY KEY(city_id)
);