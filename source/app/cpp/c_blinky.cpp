#include "rdno_blinky/c_blinky.h"
#include "rdno_core/c_linear_allocator.h"
#include "rdno_core/c_dio.h"
#include "rdno_core/c_timer.h"
#include "rdno_core/c_serial.h"

using namespace ncore;

#ifdef TARGET_ESP32

// Initialize the system
void setup()
{
    // This is where you would set up your hardware, peripherals, etc.
    ncore::nserial::Begin(ncore::nserial::nbaud::Rate115200);  // Initialize serial communication at 115200 baud
    ncore::npin::SetPinMode(2, ncore::npin::ModeOutput);  // Set the LED pin as output
}

// Main loop of the application
void loop()
{
    nserial::Println("Blinking LED...");

    // This is where you would put your main logic, such as blinking an LED.

    ncore::npin::WritePin(2, ncore::npin::High);  // Turn on the LED
    ncore::ntimer::Delay(1000);                   // Wait for 1 second
    ncore::npin::WritePin(2, ncore::npin::Low);   // Turn on the LED
    ncore::ntimer::Delay(1000);                   // Wait for 1 second
}

#else 

int main()
{
    // This is where you would set up your hardware, peripherals, etc.
    ncore::nserial::Begin(ncore::nserial::nbaud::Rate115200);  // Initialize serial communication at 115200 baud
    ncore::npin::SetPinMode(2, ncore::npin::ModeOutput);  // Set the LED pin as output

    while (true)
    {
        nserial::Println("Blinking LED...");

        // This is where you would put your main logic, such as blinking an LED.

        ncore::npin::WritePin(2, ncore::npin::High);  // Turn on the LED
        ncore::ntimer::Delay(1000);                   // Wait for 1 second
        ncore::npin::WritePin(2, ncore::npin::Low);   // Turn on the LED
        ncore::ntimer::Delay(1000);                   // Wait for 1 second
    }

    return 0;
}

#endif