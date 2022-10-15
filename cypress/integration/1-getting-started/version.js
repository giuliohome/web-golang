it('displays version info', () => {
    cy.visit( "http://" + Cypress.env('LB_IP') + "/version/" )
    cy.get('h1').should('have.length', 1)
    cy.get('h1').first().should('have.text', 'version')
    cy.get('p').last().should('have.text', 'Build User: giuliohome')
  })
