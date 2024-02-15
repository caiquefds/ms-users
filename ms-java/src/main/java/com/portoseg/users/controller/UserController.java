package com.portoseg.users.controller;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import javax.validation.Valid;

import com.portoseg.users.model.dto.response.GetUserByIdResponse;
import org.apache.commons.lang3.StringUtils;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PatchMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.bind.annotation.RestController;

import com.portoseg.users.model.dto.request.CreateUserRequest;
import com.portoseg.users.model.dto.request.UpdateUserRequest;
import com.portoseg.users.service.UserService;

import lombok.RequiredArgsConstructor;

@RestController
@RequiredArgsConstructor
@RequestMapping("/v1/user")
public class UserController {
    private final UserService userService;

    @PostMapping
    @ResponseStatus(HttpStatus.ACCEPTED)
    public void createUser(@RequestBody @Valid CreateUserRequest createUserRequest, HttpServletRequest request, HttpServletResponse response) {
        String userId = userService.createUser(createUserRequest);
        String url = request.getRequestURL().toString();
        String location = StringUtils.join(url, "/", userId);

        response.addHeader("url", location);
    }

    @GetMapping("/{id}")
    @ResponseStatus(HttpStatus.OK)
    public GetUserByIdResponse getUserById(@PathVariable String id) {
        return userService.findUserById(id);
    }

    @PatchMapping("/{id}")
    @ResponseStatus(HttpStatus.OK)
    public void updateUserById(@RequestBody @Valid UpdateUserRequest updateUserRequest, @PathVariable String id) {
        userService.updateUser(updateUserRequest, id);
    }

    @DeleteMapping("/{id}")
    @ResponseStatus(HttpStatus.NO_CONTENT)
    public void deleteUserById(@PathVariable String id) {
        userService.deleteUserById(id);
    }

}
