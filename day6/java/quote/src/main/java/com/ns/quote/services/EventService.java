package com.ns.quote.services;

import com.ns.quote.models.dto.EventDto;

public interface EventService {
    EventDto saveEvent(EventDto eventDto);
}
