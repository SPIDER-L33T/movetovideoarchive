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

Example:
```
10 * * * *   root /home/spider/movetovideoarchive /etc/motion/motion.conf
```

### Usage
```sh
./movetovideoarchive full_path_to_file_motion.conf
```
Example:
```sh
./movetovideoarchive /etc/motion/motion.conf
```
If your camera profiles are divided into different files (camera1.conf, camera2.conf, etc.) and each camera has its own directory for storing video files (individual parameter "target_dir"), then you can simply run the utility with different parameters
```sh
./movetovideoarchive /etc/motion/camera1.conf
./movetovideoarchive /etc/motion/camera2.conf
./movetovideoarchive /etc/motion/camera3.conf
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

Пример:
```
10 * * * *   root /home/spider/movetovideoarchive /etc/motion/motion.conf
```
### Использование
```sh
./movetovideoarchive full_path_to_file_motion.conf
```
Пример:
```sh
./movetovideoarchive /etc/motion/motion.conf
```
Если у Вас профили камер разбиты по разным файлам (camera1.conf, camera2.conf и т.д.) и у каждой камеры своя директория хранения видеофайлов (индивидуальный параметр "target_dir"), то можно просто запускать утилиту с разными параметрами:
```sh
./movetovideoarchive /etc/motion/camera1.conf
./movetovideoarchive /etc/motion/camera2.conf
./movetovideoarchive /etc/motion/camera3.conf
```
