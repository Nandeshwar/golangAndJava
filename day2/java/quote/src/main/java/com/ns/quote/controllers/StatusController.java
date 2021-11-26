package com.ns.quote.controllers;

import com.ns.quote.models.dto.Status;
import com.ns.quote.services.StatusService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class StatusController {
    @Autowired
    StatusService statusService;

    @GetMapping(value = "/api/status", produces = "application/json")
    public Status status() {
        return statusService.appInfo();
    }
}
