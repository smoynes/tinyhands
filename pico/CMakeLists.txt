cmake_minimum_required(VERSION 3.13)

include(pico_sdk_import.cmake)

project(PicoDemo C CXX ASM)

pico_sdk_init()

# Build a PicoDemo executable
add_executable(PicoDemo main.c)

# Link with the pico stdlib
target_link_libraries(PicoDemo pico_stdlib)

# Create map/bin/hex/uf2 files
pico_add_extra_outputs(PicoDemo)

# Enable USB serial
pico_enable_stdio_usb(PicoDemo 1)
pico_enable_stdio_uart(PicoDemo 1)
