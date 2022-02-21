package com.ns.quote.controllers;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class WelcomeController {

    @GetMapping(value = "/", produces = "application/json")
    public String Welcome() {
        return "Welcome to the world of spring boot - java";
    }
}
