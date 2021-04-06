import React from 'react';
import { act, fireEvent } from '@testing-library/react';
import ReactDOM from 'react-dom';
import Login from "./Login";
import { backendService } from '../service/backendService';



describe('Login component tests', () => {
  let container: HTMLDivElement;
  let emailInput: Element;
  let passwordInput: Element;
  let submitButton: Element;
  const loginServiceSpy = jest.spyOn(backendService, 'login')

  beforeEach(() => {
    container = document.createElement('div');
    document.body.appendChild(container);
    ReactDOM.render(<Login/>, container);

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
