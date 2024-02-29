package com.portoseg.users.model.dto.response;

import com.fasterxml.jackson.databind.PropertyNamingStrategies;
import com.fasterxml.jackson.databind.annotation.JsonNaming;
import com.portoseg.users.enumeration.StatusType;
import com.portoseg.users.model.domain.UserDomain;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;

@Data
@Builder
@AllArgsConstructor
@JsonNaming(PropertyNamingStrategies.SnakeCaseStrategy.class)
public class GetUserByIdResponse {

    private String userName;

    private String email;

    private String password;

    private StatusType status;

    public static GetUserByIdResponse valueOf(UserDomain userDomain){
       return GetUserByIdResponse.builder()
                .userName(userDomain.getUsername())
                .email(userDomain.getEmail())
                .password(userDomain.getPassword())
                .status(userDomain.getStatus())
                .build();
    }
}
