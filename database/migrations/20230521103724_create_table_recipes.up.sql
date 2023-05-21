CREATE TABLE IF NOT EXISTS recipes (
                                     id int(11) unsigned NOT NULL AUTO_INCREMENT,
                                     menu_id int(11) unsigned NOT NULL,
                                     ingredient_id int(11) unsigned NOT NULL,
                                     qty varchar(255) not null,
                                     created_at datetime DEFAULT CURRENT_TIMESTAMP,
                                     updated_at datetime DEFAULT CURRENT_TIMESTAMP,
                                     PRIMARY KEY (id)
) ENGINE=InnoDB;