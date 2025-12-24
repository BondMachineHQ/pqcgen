from sympy import *
from qiskit import QuantumCircuit
from qiskit.circuit import Parameter, ParameterVector
from qiskit_symb.quantum_info import Operator
from IPython import display

rx_1_1 = Parameter('rx_1_1')
rx_1_2 = Parameter('rx_1_2')
rx_1_3 = Parameter('rx_1_3')
rx_1_4 = Parameter('rx_1_4')
rz_1_1 = Parameter('rz_1_1')
rz_1_2 = Parameter('rz_1_2')
rz_1_3 = Parameter('rz_1_3')
rz_1_4 = Parameter('rz_1_4')

pqc = QuantumCircuit(4)
pqc.rx(rx_1_1, 0)
pqc.rx(rx_1_2, 1)
pqc.rx(rx_1_3, 2)
pqc.rx(rx_1_4, 3)
pqc.rz(rz_1_1, 0)
pqc.rz(rz_1_2, 1)
pqc.rz(rz_1_3, 2)
pqc.rz(rz_1_4, 3)

# pqc.draw('mpl')

op = Operator(pqc)
sm = op.to_sympy()
