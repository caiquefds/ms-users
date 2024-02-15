package com.portoseg.users.model.dto.response;

import com.amazonaws.services.dynamodbv2.model.Get;
import com.portoseg.users.enumeration.StatusType;
import com.portoseg.users.model.domain.UserDomain;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;

@Data
@Builder
@AllArgsConstructor
public class GetUserByIdResponse {

    private String username;

    private String email;

    private String password;

    private StatusType status;

    public static GetUserByIdResponse valueOf(UserDomain userDomain){
       return GetUserByIdResponse.builder()
                .username(userDomain.getUsername())
                .email(userDomain.getEmail())
                .password(userDomain.getPassword())
                .status(userDomain.getStatus())
                .build();
    }
}
