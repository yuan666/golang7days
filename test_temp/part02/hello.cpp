#include <iostream>
extern "C" {
   #include "hello.h"
}


void SayHelloCPP(const char *s)
{
    std::cout<<s;
}