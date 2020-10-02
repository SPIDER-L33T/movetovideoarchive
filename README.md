# movetovideoarchive
## ENG
### Description
Utility for creating an archive of videos created by [Motion](https://github.com/Motion-Project/motion)
### Requirements
1. In camera config parameter movie_filename must be:
```
movie_filename CAM01-%Y%m%d%H%M%S
```
Where is "CAM01" - name of camera

2. Add to Cron running movetovideoarchive every day

### Usage
```sh
./movetovideoarchive full_path_to_file_motion.conf
```
Example:
```sh
./movetovideoarchive /etc/motion/motion.conf
```
## RUS
### Описание
Утилита для создания видеоархива видеофайлов, созданных демоном [Motion](https://github.com/Motion-Project/motion)
### Требования
1. В конфиге камеры параметр movie_filename должен быть:
```
movie_filename CAM01-%Y%m%d%H%M%S
```
Где "CAM01" - имя камеры

2. Добавьте в Cron запуск movetovideoarchive ежедневно
### Использование
```sh
./movetovideoarchive full_path_to_file_motion.conf
```
Пример:
```sh
./movetovideoarchive /etc/motion/motion.conf
```
