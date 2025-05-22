#include "rdno_blinky/c_blinky.h"
#include "rdno_core/c_linear_allocator.h"
#include "rdno_core/c_dio.h"
#include "rdno_core/c_timer.h"
#include "rdno_core/c_serial.h"

namespace ncore
{
    static byte           sLocalMemory[1024];  // Local memory for the linear allocator
    static linear_alloc_t gLinearAlloc;

};  // namespace ncore

using namespace ncore;

// Initialize the system
void setup()
{
    gLinearAlloc.setup(sLocalMemory, sizeof(sLocalMemory));

    // This is where you would set up your hardware, peripherals, etc.
    ncore::nserial::begin(ncore::nserial::nbaud::Rate115200);  // Initialize serial communication at 115200 baud
    ncore::npin::SetPinMode(2, ncore::npin::ModeOutput);  // Set the LED pin as output
}

// Main loop of the application
void loop()
{
    nserial::println("Blinking LED...");

    // This is where you would put your main logic, such as blinking an LED.

    ncore::npin::WritePin(2, ncore::npin::High);  // Turn on the LED
    ncore::ntimer::Delay(1000);                   // Wait for 1 second
    ncore::npin::WritePin(2, ncore::npin::Low);   // Turn on the LED
    ncore::ntimer::Delay(1000);                   // Wait for 1 second
}
