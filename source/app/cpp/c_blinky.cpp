#include "rdno_blinky/c_blinky.h"
#include "rdno_core/c_linear_allocator.h"
#include "rdno_core/c_gpio.h"
#include "rdno_core/c_timer.h"
#include "rdno_core/c_serial.h"

using namespace ncore;

#ifdef TARGET_ESP32

// Initialize the system
void setup()
{
    // This is where you would set up your hardware, peripherals, etc.
    ncore::nserial::Begin(ncore::nserial::nbaud::Rate115200);  // Initialize serial communication at 115200 baud
    ncore::ngpio::set_pinmode(2, ncore::ngpio::ModeOutput);    // Set the LED pin as output
}

// Main loop of the application
void loop()
{
    nserial::Println("Blinking LED...");

    // This is where you would put your main logic, such as blinking an LED.

    ncore::ngpio::write_digital(2, ncore::ngpio::High);  // Turn on the LED
    ncore::ntimer::delay(1000);                          // Wait for 1 second
    ncore::ngpio::write_digital(2, ncore::ngpio::Low);   // Turn on the LED
    ncore::ntimer::delay(1000);                          // Wait for 1 second
}

#else

int main()
{
    // This is where you would set up your hardware, peripherals, etc.
    ncore::nserial::begin(ncore::nbaud::Rate115200);         // Initialize serial communication at 115200 baud
    ncore::ngpio::set_pinmode(2, ncore::ngpio::ModeOutput);  // Set the LED pin as output

    while (true)
    {
        nserial::println("Blinking LED...");

        // This is where you would put your main logic, such as blinking an LED.

        ncore::ngpio::write_digital(2, ncore::ngpio::High);  // Turn on the LED
        ncore::ntimer::delay(1000);                          // Wait for 1 second
        ncore::ngpio::write_digital(2, ncore::ngpio::Low);   // Turn on the LED
        ncore::ntimer::delay(1000);                          // Wait for 1 second
    }

    return 0;
}

#endif