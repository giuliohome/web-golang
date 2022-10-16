it('displays version info', () => {
    cy.visit( "http://" + Cypress.env('LB_IP') + "/version/" )
    cy.get('h1').should('have.length', 1)
    cy.get('h1').first().should('have.text', 'version')
    cy.get('p').last()
        .should(($paragraph) => {
            const versText = $paragraph.text()
            expect(versText, 'Version').to.contain('Version: ')
        })
  })
  it('displays build info', () => {
    cy.visit( "http://" + Cypress.env('LB_IP') + "/build/" )
    cy.get('h1').should('have.length', 1)
    cy.get('h1').first().should('have.text', 'build')
    cy.get('p').last().should('have.text', 'Build User: giuliohome')
  })
