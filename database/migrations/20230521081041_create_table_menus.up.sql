CREATE TABLE IF NOT EXISTS menus (
                                           id int(11) unsigned NOT NULL AUTO_INCREMENT,
                                           name varchar(255) NOT NULL,
                                           category_id int(11) unsigned NOT NULL,
                                           created_at datetime DEFAULT CURRENT_TIMESTAMP,
                                           updated_at datetime DEFAULT CURRENT_TIMESTAMP,
                                           PRIMARY KEY (id)
) ENGINE=InnoDB;