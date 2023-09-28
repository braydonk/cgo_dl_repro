#include "lib.h"
#include <stdio.h>
#include <dlfcn.h>

int main() {
	void *handle;

	handle = dlopen("./liblib.so", RTLD_LAZY);
	printf("hi :)\n");
	int x = get42();
	printf("%d\n", x);

	return 0;
}
