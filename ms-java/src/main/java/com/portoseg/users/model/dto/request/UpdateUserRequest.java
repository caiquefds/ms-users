package com.portoseg.users.model.dto.request;

import com.portoseg.users.enumeration.StatusType;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class UpdateUserRequest {

    private String username;

    private String email;

    private String password;

    private StatusType status;
}