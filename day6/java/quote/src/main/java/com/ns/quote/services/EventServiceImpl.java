package com.ns.quote.services;

import com.ns.quote.models.dto.EventDto;
import org.springframework.stereotype.Service;

@Service
public class EventServiceImpl implements EventService {
    @Override
    public EventDto saveEvent(EventDto eventDto) {
        eventDto.setId(100);
        return eventDto;
    }
}
