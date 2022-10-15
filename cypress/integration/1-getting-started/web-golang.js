/// <reference types="cypress" />

// Welcome to Cypress!
//
// This spec file contains a variety of sample tests
// for a todo list app that are designed to demonstrate
// the power of writing tests in Cypress.
//
// To learn more about how Cypress works and
// what makes it such an awesome testing tool,
// please read our getting started guide:
// https://on.cypress.io/introduction-to-cypress

describe('golang web app', () => {
  beforeEach(() => {
    // Cypress starts out with a blank slate for each test
    // so we must tell it to visit our website with the `cy.visit()` command.
    // Since we want to visit the same URL at the start of all our tests,
    // we include it in our beforeEach function so that it runs before each test
    cy.visit( "http://" + Cypress.env('LB_IP') + "/view/a1" )
  })

  it('displays one h1 title by default', () => {
    cy.get('h1').should('have.length', 1)

    cy.get('h1').first().should('have.text', 'a1')
    cy.get('div').last().should('have.text', 'web app in golang tested OK! ;-)')
  })





  it('can edit item', () => {
    // First, let's click the "Clear completed" button
    // `contains` is actually serving two purposes here.
    // First, it's ensuring that the button exists within the dom.
    // This button only appears when at least one task is checked
    // so this command is implicitly verifying that it does exist.
    // Second, it selects the button so we can click it.
    cy.contains('edit').click()

    // Then we can make sure that there is only one element
    // in the list and our element does not exist
    cy.get('textarea')
      .should('have.length', 1)
      .should('have.text', 'web app in golang tested OK! ;-)')

    // Finally, make sure that the clear button no longer exists.
    cy.contains('edit').should('not.exist')
  })

  
})
