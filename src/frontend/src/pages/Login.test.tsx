import React from 'react';
import { act, fireEvent } from '@testing-library/react';
import ReactDOM from 'react-dom';
import Login from "./Login";
import { backendService } from '../service/backendService';
import { AuthContext } from '../contexts/AuthContext';
import { Router } from 'react-router-dom';

describe('Login component tests', () => {
  let container: HTMLDivElement;
  let emailInput: Element;
  let passwordInput: Element;
  let submitButton: Element;
  const loginServiceSpy = jest.spyOn(backendService, 'login')
  const mockSetUser = jest.fn();
  const mockHistory = { push: jest.fn(), location: {}, listen: jest.fn() };

  beforeEach(() => {
    container = document.createElement('div');
    document.body.appendChild(container);
    ReactDOM.render(
      <Router history={mockHistory as any}>
        <AuthContext.Provider value={{user: 'mockUser', setUser: mockSetUser}}>
          <Login/>
        </AuthContext.Provider>
      </Router>, container);
    emailInput = container.querySelector("[id='email']") as Element
    passwordInput = container.querySelector("[id='password']") as Element
    submitButton = container.querySelector("[type='submit']") as Element
  })

  afterEach(() => {
    document.body.removeChild(container);
    container.remove();
  })

  it('Renders the Login component correctly', () => {
    expect(emailInput).toBeInTheDocument();
    expect(passwordInput).toBeInTheDocument();
    expect(submitButton).toBeInTheDocument();
  })

  it('Passes credentials correctly',  async () => {
    await act(async () => {
      fireEvent.change(emailInput, {target: {value: 'email@email.com'}});
      fireEvent.change(passwordInput, {target: {value: 'password'}});
    })
    await act(async () => {
      fireEvent.click(submitButton);
    })

    expect(loginServiceSpy).toBeCalledWith('email@email.com', 'password');
  })

  it('redirects the page to / if login successful', async() => {
    const response: any = {
      headers: {
        "csrf-token": "mockToken"
      },
      data: {
        id: 1,
        email: "email@email.com"
      }
    }
    loginServiceSpy.mockResolvedValue(response);
    await act(async () => {
      fireEvent.change(emailInput, {target: {value: 'email@email.com'}});
      fireEvent.change(passwordInput, {target: {value: 'password'}});
    })
    await act(async () => {
      fireEvent.click(submitButton as Element);
    })

    expect(mockHistory.push).toHaveBeenCalledTimes(1);
    expect(mockHistory.push).toHaveBeenCalledWith({pathname: "/"});
  })

  it('sets an errorMsg if invalid credentials are entered',  async () => {
    let errorResponse = new Error() as any;
    errorResponse.response = {
      data: {
        error: "invalid email"
      }
    }
    loginServiceSpy.mockRejectedValue(errorResponse);
    await act(async () => {
      fireEvent.change(emailInput, {target: {value: 'bad@email.com'}});
      fireEvent.change(passwordInput, {target: {value: 'password'}});
    })
    await act(async () => {
      fireEvent.click(submitButton as Element);
    })
    const errorMsg = container.querySelector("[id='errorMsg']")

    expect(errorMsg).toHaveTextContent("invalid email");
  })
})
