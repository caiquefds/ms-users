package com.portoseg.users.model.dto.request;

import javax.validation.constraints.Email;
import javax.validation.constraints.NotNull;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class CreateUserRequest {

    @NotNull(message = "400.002")
    private String username;

    @Email(message = "400.003")
    @NotNull(message = "400.002")
    private String email;

    @NotNull(message = "400.002")
    private String password;

}