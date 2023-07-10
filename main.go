package main

// #include <string.h>
// #include <stdlib.h>
// typedef struct{
//  int id;
//  char *name;
//}output_data_t;
// typedef int (*intFunc) (output_data_t *d);
// int
// bridge_int_func(intFunc f, output_data_t *d)
// {
//		return f(d);
// }
//
// int fill_d(output_data_t *d)
// {
//	d->id = 5;
//      d->name = calloc(20, sizeof(char));
//	strcpy(d->name, "test\0");
//	return 0;
// }
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	f := C.intFunc(C.fill_d)
	data := C.output_data_t{}
	C.bridge_int_func(f, &data)
	fmt.Println(int(data.id))
	fmt.Println(C.GoString(data.name))
	C.free(unsafe.Pointer(data.name))

}
