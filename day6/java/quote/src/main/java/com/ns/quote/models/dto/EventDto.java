package com.ns.quote.models.dto;

import lombok.Getter;
import lombok.Setter;

import javax.validation.constraints.Max;
import javax.validation.constraints.Min;
import javax.validation.constraints.NotNull;
import java.time.ZonedDateTime;

@Getter
@Setter
public class EventDto {
    private int id;
    @NotNull
    @Min(1)
    @Max(31)
    private int day;
    private int month;
    private int year;
    private String title;
    private String description;
    private ZonedDateTime createdAt;
    private ZonedDateTime updatedAt;
}
