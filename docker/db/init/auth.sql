create database auth;

grant all privileges on auth.*
    TO 'wallche_auth'@'%' identified by 'wallche_auth_pass134292725';

flush privileges;