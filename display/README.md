# Display Hat for RPI Zero W

Configure Raspberry Pi to enable interface for SPI:

```console
$ sudo raspi-config
```

Install system packages:

```
$ sudo apt install python3-numpy python3-rpi.gpio python3-pil
```

```console
$ python3 -m venv --system-site-packages env
$ ./env/bin/pip install displayhatmini
$ ./env/bin/python3 display.py
```
